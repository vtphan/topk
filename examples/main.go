package main
import (
   "fmt"
   "time"
   "math/rand"
   "github.com/vtphan/topk"
)

type Data struct{}

func (d *Data) IsBetter(id1, id2 int) bool {
   return id1 < id2
}


func main() {
   r := rand.New(rand.NewSource(time.Now().UnixNano()))
   K := 5
   N := 15

   h := topk.NewHeap(&Data{}, K)

   fmt.Println("Inserting random numbers into the queue")
   for i:=0; i<N; i++ {
      x := r.Intn(1000)
      h.Push(x)
      fmt.Print("Insert ", x)
      h.Show()
   }
   fmt.Print("The best ", K, " elements are: ", h.Get(), "\n")
   fmt.Println("Popping each element at a time:")
   for h.Size() > 0 {
      fmt.Print(h.Pop(),",")
   }
}