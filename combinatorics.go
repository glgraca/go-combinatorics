//Package combinatorics provides functions for generating combinations (n-choose-k, permutations, and derangements) of arrays of any kind of value.
package combinatorics

func factoradics(n int, ch chan []int, indices []int) {
  if n==1 {
    indices=append(indices, 0)
    ch<-indices
  } else {
    for i:=0; i<n; i++ {
      factoradics(n-1, ch, append(indices, i))
    }
  }
}

//Factoradics will count in factoradic base.
//It returns an array of []int with the digits of each number.
func Factoradics(n int) <-chan []int {
  ch:=make(chan []int)
  go func() {
    factoradics(n, ch, nil)
    close(ch)
  }()
  return ch
}

//Permutations will return a channel that will produce arrays with permutations of the original set.
func Permutations(vals []interface{}) <- chan []interface{} {
  ch:=make(chan []interface{})
  fact:=Factoradics(len(vals))

  go func() {
    for r:=range fact {
      tmp:=make([]interface{}, len(r))
      copy(tmp, vals)
      permutation:=make([]interface{}, len(r))
      for i:=0; i<len(r); i++ {
        v:=tmp[r[i]]
        tmp=append(tmp[:r[i]], tmp[r[i]+1:]...)
        permutation[i]=v
      }
      ch<-permutation
    }
    close(ch)
  }()

  return ch
}

func choose(n int, k int, vals []interface{}, ch chan []interface{}, indices []int) {
  if k>0 {
    for m:=k; m<=n; m++ {
      choose(m-1, k-1, vals, ch, append([]int{m},indices...)); 
      if k==1 {
        line:=make([]interface{},len(indices)+1)
        line[0]=vals[m-1]
        for i:=0; i<len(indices); i++ {
          line[i+1]=vals[indices[i]-1]
        }
        ch<-line
      }
    }
  } 
}


//Choose returns a channel that produces arrays of subsets of the original set in groups of k elements.
func Choose(k int, vals []interface{}) <-chan []interface{} {
  ch:=make(chan []interface{});  
  go func() {
    choose(len(vals), k, vals, ch, nil);
    close(ch)
  }()
  return ch
}


//Contains verifies if an element val is in the set vals
func Contains(vals []interface{}, val interface{}) bool {
  for _,v := range vals {
    if val == v {
      return true
    }
  }
  return false
}

func derangements(n int, vals []interface{}, ch chan []interface{}, derangement []interface{}) {
  if n==len(vals) {
    ch<-derangement
  } else {
    columns:=make([]interface{}, len(vals))
    copy(columns, vals)
    columns=append(columns[:n], columns[n+1:]...)
    for _,column:=range columns {
      if !Contains(derangement, column) {
        d:=append(derangement, column)
        derangements(n+1, vals, ch, d)
      }
    }
  }
}

//Derangements produces permutations of vals in which no element is in its original position
func Derangements(vals []interface{}) <- chan []interface{} {
  ch:=make(chan []interface{})
  go func() {
    derangements(0, vals, ch, nil)
    close(ch)
  }()
  return ch
}