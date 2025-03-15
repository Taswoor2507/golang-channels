// // // // // // // // // // // package main

// // // // // // // // // // // import (
// // // // // // // // // // // 	"fmt"
// // // // // // // // // // // 	"time"
// // // // // // // // // // // )

// // // // // // // // // // // func serveFood(table int) {
// // // // // // // // // // // 	fmt.Println("Serving food to table", table)
// // // // // // // // // // // }
// // // // // // // // // // // func main() {
// // // // // // // // // // // 	for i := 1; i < 1000; i++ {
// // // // // // // // // // // 		go serveFood(i)
// // // // // // // // // // // 	}
// // // // // // // // // // // 	time.Sleep(1 * time.Second)
// // // // // // // // // // // 	fmt.Println("Alltable served successfully")
// // // // // // // // // // // }

// // // // // // // // // // // package main

// // // // // // // // // // // import (
// // // // // // // // // // // 	"fmt"
// // // // // // // // // // // 	"reflect"
// // // // // // // // // // // 	"sync"
// // // // // // // // // // // )

// // // // // // // // // // // func serveFood(table int, wg *sync.WaitGroup) {
// // // // // // // // // // // 	defer wg.Done()
// // // // // // // // // // // 	fmt.Printf("%+v\n", wg)         // Print WaitGroup internal fields
// // // // // // // // // // // 	fmt.Println(reflect.TypeOf(wg)) // Check the actual struct type
// // // // // // // // // // // 	fmt.Println("Serving food to table----->", table, *wg)
// // // // // // // // // // // }
// // // // // // // // // // // func main() {
// // // // // // // // // // // 	var wg sync.WaitGroup
// // // // // // // // // // // 	for i := 1; i < 50000; i++ {
// // // // // // // // // // // 		wg.Add(1)
// // // // // // // // // // // 		go serveFood(i, &wg)

// // // // // // // // // // // 	}
// // // // // // // // // // // 	wg.Wait()
// // // // // // // // // // // 	fmt.Println("All table served successfully")

// // // // // // // // // // // }

// // // // // // // // // // // defer
// // // // // // // // // // // package main

// // // // // // // // // // // import (
// // // // // // // // // // // 	"fmt"
// // // // // // // // // // // )

// // // // // // // // // // // func main() {
// // // // // // // // // // // 	fmt.Println("Hello")
// // // // // // // // // // // 	defer fmt.Println("Bye")
// // // // // // // // // // // 	defer fmt.Println("World")
// // // // // // // // // // // }

// // // // // // // // // // //

// // // // // // // // // // // func serveDrinks(table int, wg *sync.WaitGroup) {
// // // // // // // // // // // 	defer wg.Done() // Ensure Done is called even if the function panics
// // // // // // // // // // // 	fmt.Println("Serving drinks to table", table)
// // // // // // // // // // // }

// // // // // // // // // // // func serveFood(table int, wg *sync.WaitGroup) {
// // // // // // // // // // // 	defer wg.Done() // Ensure Done is called even if the function panics
// // // // // // // // // // // 	fmt.Println("Serving food to table", table)
// // // // // // // // // // // }

// // // // // // // // // // // func main() {
// // // // // // // // // // // 	var wg sync.WaitGroup
// // // // // // // // // // // 	var wg2 sync.WaitGroup
// // // // // // // // // // // 	// Launch 5 goroutines for serveDrinks
// // // // // // // // // // // 	for i := 1; i <= 5; i++ {
// // // // // // // // // // // 		wg.Add(1) // Add 1 for each goroutine
// // // // // // // // // // // 		go serveDrinks(i, &wg)
// // // // // // // // // // // 	}

// // // // // // // // // // // 	// Launch 5 goroutines for serveFood
// // // // // // // // // // // 	for i := 1; i <= 5; i++ {
// // // // // // // // // // // 		wg2.Add(1) // Add 1 for each goroutine
// // // // // // // // // // // 		go serveFood(i, &wg2)
// // // // // // // // // // // 	}

// // // // // // // // // // // 	// Wait for all goroutines to finish
// // // // // // // // // // // 	wg.Wait()
// // // // // // // // // // // 	fmt.Println("All tables served successfully")
// // // // // // // // // // // }

// // // // // // // // // // // https://go.dev/tour/concurrency/3

// // // // // // // // // // package main

// // // // // // // // // // import (
// // // // // // // // // // 	"fmt"
// // // // // // // // // // 	"time"
// // // // // // // // // // )

// // // // // // // // // // func reciever(ch chan int) {
// // // // // // // // // // 	fmt.Println("Reciever: ", <-ch)
// // // // // // // // // // 	// Blocks until a value is sent to ch

// // // // // // // // // // }
// // // // // // // // // // func main() {
// // // // // // // // // // 	fmt.Println("Channels in golang ")
// // // // // // // // // // 	time.Sleep(time.Second * 5)
// // // // // // // // // // 	ch1 := make(chan int)
// // // // // // // // // // 	go reciever(ch1) // Start a new goroutine to handle reciever function
// // // // // // // // // // 	ch1 <- 1
// // // // // // // // // // 	// fmt.Println(<-ch1) // Prints: 1
// // // // // // // // // // 	fmt.Println("hello wrold ")
// // // // // // // // // // }

// // // // // // // // // package main

// // // // // // // // // import (
// // // // // // // // // 	"fmt"
// // // // // // // // // 	"io"
// // // // // // // // // 	"net/http"
// // // // // // // // // )

// // // // // // // // // func fetchURL(url string, ch chan<- string) {
// // // // // // // // // 	resp, err := http.Get(url)
// // // // // // // // // 	if err != nil {
// // // // // // // // // 		ch <- fmt.Sprintf("Error fetching %s: %v", url, err)
// // // // // // // // // 		return
// // // // // // // // // 	}
// // // // // // // // // 	defer resp.Body.Close() // Ensure response body is closed after reading

// // // // // // // // // 	// Read the response body
// // // // // // // // // 	body, err := io.ReadAll(resp.Body)
// // // // // // // // // 	if err != nil {
// // // // // // // // // 		ch <- fmt.Sprintf("Error reading body of %s: %v", url, err)
// // // // // // // // // 		return
// // // // // // // // // 	}
// // // // // // // // // https://github.com/LeHaGiaBao/Eleven-Golang-Projects/tree/master/Email%20Verifier%20Tool%20With%20Golang
// // // // // // // // // https://github.com/kkdai/project52
// // // // // // // // // https://github.com/avelino/awesome-go
// // // // // // // // // 	// Send response status and first 100 characters of the response body
// // // // // // // // // 	ch <- fmt.Sprintf("Fetched %s with status: %s\nResponse Body (First 100 chars): %s\n",
// // // // // // // // // 		url, resp.Status, string(body[:10000]))
// // // // // // // // // }

// // // // // // // // // func main() {
// // // // // // // // // 	urls := []string{
// // // // // // // // // 		"https://golang.org",
// // // // // // // // // 		"https://google.com",
// // // // // // // // // 		"https://github.com",
// // // // // // // // // 	}

// // // // // // // // // 	ch := make(chan string)
// // // // // // // // // 	for _, url := range urls {
// // // // // // // // // 		go fetchURL(url, ch)
// // // // // // // // // 	}

// // // // // // // // // 	// Receive and print responses
// // // // // // // // // 	for range urls {
// // // // // // // // // 		fmt.Println(<-ch)
// // // // // // // // // 	}
// // // // // // // // // }

// // // // // // // // // package main

// // // // // // // // // import (
// // // // // // // // // 	"fmt"
// // // // // // // // // 	"time"
// // // // // // // // // )

// // // // // // // // // sending
// // // // // // // // // func processNum(numchan chan int) {
// // // // // // // // // 	for num := range numchan {

// // // // // // // // // 		fmt.Println("processing number ..", num)
// // // // // // // // // 		time.Sleep(time.Second)
// // // // // // // // // 	}
// // // // // // // // // }

// // // // // // // // // recieving
// // // // // // // // //
// // // // // // // // //	func sum(result chan int, num1 int, num2 int) {
// // // // // // // // //		resultOf := num1 + num2
// // // // // // // // //		result <- resultOf
// // // // // // // // //	}
// // // // // // // // //
// // // // // // // // // goroutine synccronizer
// // // // // // // // //
// // // // // // // // //	func task(done chan bool) {
// // // // // // // // //		defer func() { done <- true }()
// // // // // // // // //		fmt.Println("processing task...")
// // // // // // // // //	}

// // // // // // // // // func emailSender(emailChan <-chan string, done chan<- bool) {
// // // // // // // // // 	defer func() { done <- true }()
// // // // // // // // // 	for email := range emailChan {
// // // // // // // // // 		fmt.Println("Sending email to ", email)
// // // // // // // // // 		time.Sleep(time.Second)
// // // // // // // // // 	}
// // // // // // // // // }

// // // // // // // // // func main() {
// // // // // // // // // 	chan1 := make(chan int)
// // // // // // // // // 	chan2 := make(chan string)

// // // // // // // // // 	go func() {
// // // // // // // // // 		chan1 <- 10
// // // // // // // // // 	}()
// // // // // // // // // 	go func() {
// // // // // // // // // 		chan2 <- "pong"
// // // // // // // // // 	}()

// // // // // // // // // 	// combination
// // // // // // // // // 	for i := 0; i < 2; i++ {
// // // // // // // // // 		select {
// // // // // // // // // 		case chan1Value := <-chan1:
// // // // // // // // // 			fmt.Println("received data from channnel chan1", chan1Value)
// // // // // // // // // 		case chan2Value := <-chan2:
// // // // // // // // // 			fmt.Println("received data from channnel chan2", chan2Value)
// // // // // // // // // 		}
// // // // // // // // // 	}

// // // // // // // // // 	// emailChan := make(chan string, 100) //buffered channel
// // // // // // // // // 	// done := make(chan bool)
// // // // // // // // // 	// go emailSender(emailChan, done)
// // // // // // // // // 	// for i := 0; i < 5; i++ {
// // // // // // // // // 	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
// // // // // // // // // 	// }
// // // // // // // // // 	// close(emailChan)
// // // // // // // // // 	// fmt.Println("done sending ")
// // // // // // // // // 	// <-done

// // // // // // // // // 	// emailChan <- "one@example.com"
// // // // // // // // // 	// emailChan <- "two@example.com"
// // // // // // // // // 	// emailChan <- "three@example.com"
// // // // // // // // // 	// fmt.Println("chan data", <-emailChan)
// // // // // // // // // 	// fmt.Println("chan data", <-emailChan)
// // // // // // // // // 	// fmt.Println("chan data", <-emailChan)

// // // // // // // // // 	// done := make(chan bool)
// // // // // // // // // 	// go task(done)
// // // // // // // // // 	// <-done //block jab tak data send nahi hota
// // // // // // // // // 	// fmt.Println("end")
// // // // // // // // // 	// result := make(chan int)
// // // // // // // // // 	// go sum(result, 4, 5)
// // // // // // // // // 	// res := <-result
// // // // // // // // // 	// fmt.Println(res)

// // // // // // // // // 	// numchan := make(chan int)
// // // // // // // // // 	// go processNum(numchan)
// // // // // // // // // 	// for {
// // // // // // // // // 	// 	numchan <- rand.Intn(100)
// // // // // // // // // 	// }
// // // // // // // // // 	// numchan <- 10

// // // // // // // // // 	// time.Sleep(time.Second)

// // // // // // // // // 	// messageChan := make(chan string)
// // // // // // // // // 	// messageChan <- "Hello word " //blocking
// // // // // // // // // 	// msg := <-messageChan
// // // // // // // // // 	// fmt.Println(msg)
// // // // // // // // // }

// // // // // // // // // mutex in golang
// // // // // // // // // package main

// // // // // // // // // type Post struct {
// // // // // // // // // 	views int
// // // // // // // // // 	mx    sync.Mutex
// // // // // // // // // }

// // // // // // // // // func (P *Post) incViews(wg *sync.WaitGroup) {
// // // // // // // // // 	// lock the mutex
// // // // // // // // // 	// p.mu.Lock()
// // // // // // // // // 	// defer p.mu.Unlock()
// // // // // // // // // 	// p.views++
// // // // // // // // // 	defer wg.Done()
// // // // // // // // // 	P.mx.Lock()
// // // // // // // // // 	P.views++
// // // // // // // // // 	fmt.Println(P.views)
// // // // // // // // // 	defer P.mx.Unlock()

// // // // // // // // // }
// // // // // // // // // func main() {
// // // // // // // // // 	post := &Post{}
// // // // // // // // // 	var wg sync.WaitGroup
// // // // // // // // // 	for i := 0; i < 357; i++ {
// // // // // // // // // 		wg.Add(1)
// // // // // // // // // 		go post.incViews(&wg)
// // // // // // // // // 	}
// // // // // // // // // 	wg.Wait()
// // // // // // // // // 	fmt.Println("Total views: ", post.views)
// // // // // // // // // }

// // // // // // // // // files in golang
// // // // // // // // // ________________________________________________
// // // // // // // // // package main

// // // // // // // // // import (
// // // // // // // // // 	"fmt"
// // // // // // // // // 	"os"
// // // // // // // // // )

// // // // // // // // // func main() {
// // // // // // // // // 	f, err := os.Open("example.txt")

// // // // // // // // // 	if err != nil {
// // // // // // // // // 		panic(err)
// // // // // // // // // 	}

// // // // // // // // // 	fileInfo, err := f.Stat()

// // // // // // // // // 	if err != nil {
// // // // // // // // // 		panic(err)
// // // // // // // // // 	}

// // // // // // // // // 	fmt.Println("file name", fileInfo.Name())
// // // // // // // // // 	fmt.Println("file size", fileInfo.Size())
// // // // // // // // // 	fmt.Println("last modified time", fileInfo.ModTime())
// // // // // // // // // 	fmt.Println("Is folder ?", fileInfo.IsDir())
// // // // // // // // // 	fmt.Println("permissions", fileInfo.Mode())

// // // // // // // // // }

// // // // // // // // // read file
// // // // // // // // // func main() {
// // // // // // // // // 	f, err := os.Open("example.txt")

// // // // // // // // // 	if err != nil {
// // // // // // // // // 		panic(err)
// // // // // // // // // 	}

// // // // // // // // // 	fileInfo, err := f.Stat()

// // // // // // // // // 	if err != nil {
// // // // // // // // // 		panic(err)
// // // // // // // // // 	}

// // // // // // // // // 	defer f.Close()

// // // // // // // // // 	//   creae a buffer
// // // // // // // // // 	b := make([]byte, fileInfo.Size())
// // // // // // // // // 	d, err := f.Read(b)

// // // // // // // // // 	if err != nil {
// // // // // // // // // 		panic(err)
// // // // // // // // // 	}

// // // // // // // // // 	fmt.Println("value of d and b ", d, "\n", string(b))
// // // // // // // // // 	// for i := 0; i < len(b); i++ {
// // // // // // // // // 	// 	fmt.Println("data : ", string(b[i]))
// // // // // // // // // 	// }

// // // // // // // // // 	fmt.Println("file read success")
// // // // // // // // // }

// // // // // // // // // method two
// // // // // // // // // func main() {
// // // // // // // // // 	data, err := os.ReadFile("example.txt")
// // // // // // // // // 	if err != nil {
// // // // // // // // // 		panic(err)
// // // // // // // // // 	}
// // // // // // // // // 	fmt.Println(string(data))
// // // // // // // // // }

// // // // // // // // // func main() {
// // // // // // // // // 	// create file
// // // // // // // // // 	file, err := os.Create("example2.txt")
// // // // // // // // // 	if err != nil {
// // // // // // // // // 		panic(err)
// // // // // // // // // 	}
// // // // // // // // // 	defer file.Close()
// // // // // // // // // 	//    add data to file
// // // // // // // // // 	file.WriteString("Hello from golang land")
// // // // // // // // // 	fmt.Println("FILE CREATED SUCCESSFULLY!")

// // // // // // // // // 	// read file
// // // // // // // // // 	f, err := os.Open("example2.txt")
// // // // // // // // // 	if err != nil {
// // // // // // // // // 		panic(err)
// // // // // // // // // 	}
// // // // // // // // // 	defer f.Close()
// // // // // // // // // 	// extract file info
// // // // // // // // // 	fileInfo, err := f.Stat()
// // // // // // // // // 	if err != nil {
// // // // // // // // // 		panic(err)
// // // // // // // // // 	}

// // // // // // // // // 	// create buffer
// // // // // // // // // 	b := make([]byte, fileInfo.Size())
// // // // // // // // // 	d, err := f.Read(b)
// // // // // // // // // 	fmt.Println("file data--> ", string(b), " <-- size ", d)

// // // // // // // // // }

// // // // // // // // // working with folders
// // // // // // // // // func main() {
// // // // // // // // // 	dir, err := os.Open(".")
// // // // // // // // // 	if err != nil {
// // // // // // // // // 		panic(err)
// // // // // // // // // 	}
// // // // // // // // // 	defer dir.Close()
// // // // // // // // // 	files, err := dir.ReadDir(-1)
// // // // // // // // // 	if err != nil {
// // // // // // // // // 		panic(err)
// // // // // // // // // 	}
// // // // // // // // // 	for _, file := range files {
// // // // // // // // // 		fmt.Println(file.Name())
// // // // // // // // // 		fmt.Println(file.IsDir())
// // // // // // // // // 		fmt.Println("---------------------------------")
// // // // // // // // // 	}
// // // // // // // // // }

// // // // // // // // // func main() {
// // // // // // // // // 	file, err := os.Create("example.txt")

// // // // // // // // // 	if err != nil {
// // // // // // // // // 		panic(err)
// // // // // // // // // 	}

// // // // // // // // // 	defer file.Close()
// // // // // // // // // 	file.WriteString("Hello , world ")

// // // // // // // // // 	// read file
// // // // // // // // // 	f, err := os.Open("example.txt")
// // // // // // // // // 	if err != nil {
// // // // // // // // // 		panic(err)
// // // // // // // // // 	}
// // // // // // // // // 	defer f.Close()
// // // // // // // // // 	fileInfo, err := f.Stat()
// // // // // // // // // 	if err != nil {
// // // // // // // // // 		panic(err)
// // // // // // // // // 	}

// // // // // // // // // 	//create byte
// // // // // // // // // 	b := make([]byte, fileInfo.Size())
// // // // // // // // // 	d, err := f.Read(b)
// // // // // // // // // 	fmt.Println("file data--> ", string(b), " <-- size ", d)

// // // // // // // // // }
// // // // // // // // // _________________________________________________________

// // // // // // // // // package main

// // // // // // // // // import (
// // // // // // // // // 	"fmt"
// // // // // // // // // 	"sync"
// // // // // // // // // )
// // // // // // // // //
// // // // // // // // // func serveFood(table int, wg *sync.WaitGroup) {
// // // // // // // // // 	defer wg.Done()
// // // // // // // // // 	fmt.Println("Serving food to table", table)
// // // // // // // // // }

// // // // // // // // //	func main() {
// // // // // // // // //		tables := 500
// // // // // // // // //		var wg sync.WaitGroup
// // // // // // // // //		wg.Add(tables)
// // // // // // // // //		for i := 0; i < tables; i++ {
// // // // // // // // //			go serveFood(i+1, &wg)
// // // // // // // // //		}
// // // // // // // // //		wg.Wait()
// // // // // // // // //		fmt.Println("All tables are seated.")
// // // // // // // // //	}
// // // // // // // // package main

// // // // // // // // import (
// // // // // // // // 	"fmt"
// // // // // // // // 	"math/rand"
// // // // // // // // 	"sync"
// // // // // // // // 	"time"
// // // // // // // // )

// // // // // // // // type Order struct {
// // // // // // // // 	ID     int
// // // // // // // // 	Status string
// // // // // // // // }

// // // // // // // // func generateOrders(count int) []*Order {
// // // // // // // // 	orders := make([]*Order, count)
// // // // // // // // 	for i := 0; i < count; i++ {
// // // // // // // // 		orders[i] = &Order{ID: i + 1, Status: "Pending"}
// // // // // // // // 	}
// // // // // // // // 	return orders
// // // // // // // // }

// // // // // // // // //process orders

// // // // // // // // func processOrders(orders []*Order) {
// // // // // // // // 	for _, order := range orders {
// // // // // // // // 		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
// // // // // // // // 		fmt.Printf("Processing order %d\n", order.ID)
// // // // // // // // 	}
// // // // // // // // }

// // // // // // // // // update order statuses
// // // // // // // // func updateOrderStatus(orders []*Order) {
// // // // // // // // 	for _, order := range orders {
// // // // // // // // 		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
// // // // // // // // 		statuses := []string{"Processing", "Shipped", "Delivered"}
// // // // // // // // 		order.Status = statuses[rand.Intn(3)]
// // // // // // // // 		fmt.Printf("Updated order %d status to %s \n", order.ID, order.Status)
// // // // // // // // 	}
// // // // // // // // }

// // // // // // // // // report order status
// // // // // // // // func reportOrderStatus(orders []*Order) {
// // // // // // // // 	for i := 0; i < 5; i++ {
// // // // // // // // 		time.Sleep((1 * time.Second))
// // // // // // // // 		fmt.Println("\n ___________________Order Status report _______________")
// // // // // // // // 		for _, order := range orders {
// // // // // // // // 			fmt.Printf("Order  %d   %s \n ", order.ID, order.Status)
// // // // // // // // 		}
// // // // // // // // 		fmt.Println("\n ________________________________________________")
// // // // // // // // 	}

// // // // // // // // }
// // // // // // // //
// // // // // // // //	func main() {
// // // // // // // //		var wg sync.WaitGroup
// // // // // // // //		wg.Add(3)
// // // // // // // //		orders := generateOrders(1000)
// // // // // // // //		go func() {
// // // // // // // //			defer wg.Done()
// // // // // // // //			processOrders(orders)
// // // // // // // //		}()
// // // // // // // //		go func() {
// // // // // // // //			defer wg.Done()
// // // // // // // //			updateOrderStatus(orders)
// // // // // // // //		}()
// // // // // // // //		go func() {
// // // // // // // //			defer wg.Done()
// // // // // // // //			reportOrderStatus(orders)
// // // // // // // //		}()
// // // // // // // //		wg.Wait()
// // // // // // // //		fmt.Println("All orders processed successfully existing..........")
// // // // // // // //	}
// // // // // // // package main

// // // // // // // import (
// // // // // // // 	"fmt"
// // // // // // // 	"math"
// // // // // // // )

// // // // // // // type Shape interface {
// // // // // // // 	Area() float64
// // // // // // // }

// // // // // // // type Rectangle struct {
// // // // // // // 	length float64
// // // // // // // 	width  float64
// // // // // // // }

// // // // // // // type Circle struct {
// // // // // // // 	radius float64
// // // // // // // }

// // // // // // // func (r Rectangle) Area() float64 {
// // // // // // // 	return r.length * r.width
// // // // // // // }

// // // // // // // func (c Circle) Area() float64 {
// // // // // // // 	return math.Pi * c.radius * c.radius
// // // // // // // }

// // // // // // // func main() {
// // // // // // // 	fmt.Println("INTERFACES in go")
// // // // // // // 	// interfaces are a way to achieve abstraction and polymorphism in go
// // // // // // // 	// an interface is a collection of method signatures that a type can implement
// // // // // // // 	// a type implements an interface by implementing its methods
// // // // // // // 	// an interface can be implemented by any type
// // // // // // // 	// an interface can be used as a type

// // // // // // // 	r := Rectangle{length: 10, width: 5}
// // // // // // // 	c := Circle{radius: 3}
// // // // // // // 	// shapes := []Shape{r, c}
// // // // // // // 	// for _, shape := range shapes {
// // // // // // // 	// 	fmt.Println("Area of shape is ", shape.Area())
// // // // // // // 	// }
// // // // // // // 	fmt.Println(r.Area())
// // // // // // // 	fmt.Println(c.Area())

// // // // // // // }

// // // // // // package main

// // // // // // import (
// // // // // // 	"fmt"
// // // // // // 	"sync"
// // // // // // )

// // // // // // type Post struct {
// // // // // // 	id    int
// // // // // // 	views int
// // // // // // 	mu    sync.Mutex
// // // // // // }

// // // // // // func (p *Post) incViews(wg *sync.WaitGroup) {
// // // // // // 	defer func() {
// // // // // // 		p.mu.Unlock()
// // // // // // 		wg.Done()
// // // // // // 	}()
// // // // // // 	p.mu.Lock()
// // // // // // 	p.views = p.views + 1
// // // // // // }

// // // // // //	func main() {
// // // // // //		wg := sync.WaitGroup{}
// // // // // //		p1 := Post{id: 1, views: 10}
// // // // // //		// p1.incViews()
// // // // // //		for i := 0; i <= 10000; i++ {
// // // // // //			wg.Add(1)
// // // // // //			go p1.incViews(&wg)
// // // // // //		}
// // // // // //		wg.Wait()
// // // // // //		fmt.Println(p1.views)
// // // // // //	}
// // // // // //
// // // // // // CONTEX WITH TIME OUT
// // // // // package main

// // // // // import (
// // // // // 	"context"
// // // // // 	"fmt"
// // // // // 	"time"
// // // // // )

// // // // // func homeWork(ctx context.Context) {
// // // // // 	count := 0
// // // // // 	for {
// // // // // 		select {
// // // // // 		case <-ctx.Done():
// // // // // 			fmt.Println("MOM SAYS  TIMES UP! CLOSE YOUR BOOKS")
// // // // // 			return
// // // // // 		default:
// // // // // 			count++
// // // // // 			fmt.Println("Homework is running now!", count)
// // // // // 		}
// // // // // 	}
// // // // // }

// // // // // func main() {
// // // // // 	// ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// // // // // 	// defer cancel()
// // // // // 	// fmt.Println(ctx, "CTX")
// // // // // 	// homeWork(ctx)

// // // // // 	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
// // // // // 	defer cancel()
// // // // // 	homeWork(ctx)

// // // // // }

// // // // // practice contex

// // // // // package main

// // // // // func orderService(ctx context.Context) {

// // // // // 	select {
// // // // // 	case <-time.After(4 * time.Second):
// // // // // 		fmt.Println("Order processing successfully!")
// // // // // 	case <-ctx.Done():
// // // // // 		fmt.Println("TIMEOUT Order processed failed!")

// // // // // 	}
// // // // // }

// // // // // func main() {
// // // // // 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

// // // // // 	defer cancel()
// // // // // 	fmt.Println("Start Order processing .....")
// // // // // 	orderService(ctx)
// // // // // }

// // // // // deadline

// // // // // func processOrder(ctx context.Context) {
// // // // // 	deadline, ok := ctx.Deadline()
// // // // // 	if !ok {
// // // // // 		fmt.Println("No deadline set")
// // // // // 		return
// // // // // 	}
// // // // // 	fmt.Println("Order processing started at ", time.Now(), "deadline", deadline)
// // // // // 	select {
// // // // // 	case <-time.After(2 * time.Second):
// // // // // 		fmt.Printf(" \n Order processing successfully at %s\n", time.Now())
// // // // // 	case <-ctx.Done():
// // // // // 		fmt.Printf("Order processing cancelled at %s\n", time.Now())
// // // // // 	}
// // // // // }

// // // // // func main() {
// // // // // 	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
// // // // // 	defer cancel()
// // // // // 	processOrder(ctx)
// // // // // }

// // // // // ________________________________________________________________________

// // // // package main

// // // // import (
// // // // 	"context"
// // // // 	"fmt"
// // // // 	"time"
// // // // )

// // // // func orderProcess(ctx context.Context, orderID string) {

// // // // 	orderId := orderID

// // // // 	select {
// // // // 	case <-time.After(3 * time.Second):
// // // // 		fmt.Printf("Order %s processed successfully\n", orderId)
// // // // 		return
// // // // 	case <-ctx.Done():
// // // // 		fmt.Printf("TIMEOUT:Order %s processing cancelled\n", orderId)
// // // // 		return

// // // // 	}

// // // // }
// // // // func main() {
// // // // 	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
// // // // 	defer cancel()

// // // // 	ctx = context.WithValue(ctx, "orderId", "GB1239KX43U789")
// // // // 	fmt.Println("Loading.......")

// // // // 	orderProcess(ctx, ctx.Value("orderId").(string))

// // // // }

// // // package main

// // // import (
// // // 	"context"
// // // 	"fmt"
// // // 	"sync"
// // // 	"time"
// // // )

// // // func route(wg *sync.WaitGroup, ctx context.Context, cancel context.CancelFunc) {
// // // 	defer wg.Done()
// // // 	select {
// // // 	case <-ctx.Done():
// // // 		fmt.Println("Order processing cancelled!")
// // // 	case <-time.After(time.Second * 3):
// // // 		fmt.Println("Order processing completed!")
// // // 		cancel()
// // // 	}
// // // }

// // // func main() {
// // // 	var wg sync.WaitGroup
// // // 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// // // 	defer cancel()

// // // 	wg.Add(1)
// // // 	go route(&wg, ctx, cancel)

// // // 	wg.Wait()
// // // }

// // package main

// // import (
// // 	"fmt"
// // 	"math/rand"
// // 	"time"
// // )

// // func main() {
// // 	src := rand.NewSource(time.Now().UnixNano())
// // 	src2 := rand.NewSource(time.Now().UnixNano() + 1)
// // 	r := rand.New(src)
// // 	r2 := rand.New(src2)
// // 	// fmt.Println(src)
// // 	num := r.Intn(2)
// // 	num2 := r2.Intn(2)
// // 	fmt.Println(num, num2)
// // 	// time.Sleep(time.Second * 5)
// // 	// fmt.Println(num)

// // }

// // // func main() {
// // // 	src := rand.NewSource(time.Now().UnixNano()) // Create a new source
// // // 	r := rand.New(src)                           // Create a new random generator

// // // 	randomNumber := r.Intn(10) // Generates a random number between 0 and 9
// // // 	fmt.Println(randomNumber)
// // // }
// // // HTML PARSER <CREATE DOM>  DOM YAY BATATI HY KI PAGE KA STRUCTURE KYA HOGA
// // // CSSOM <css object model >
// // // html + css <Render tree >
// // // rendering engine
// // // graphic card

// // // HTML PARSER
// // // CSSOM
// // // html + css <Render tree >
// // // rendering engine
// // // graphic card

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func dataSend(ch chan string) {
// 	ch <- "Hello world "
// 	fmt.Println(<-ch)
// }

// func main() {
// 	ch := make(chan string)
// 	fmt.Println("hello in the world of golang channels ...")
// 	go dataSend(ch)
// 	fmt.Println(string(<-ch))
// 	time.Sleep(time.Second * 3)
// 	fmt.Println("Program end ")

// }

package main

import (
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
)

func main() {
	data := humanize.Bytes(32323213123123123)
	data2 := humanize.Time(time.Date(2000, 8, 1, 1, 1, 1, 1, time.Now().Local().Location()))
	fmt.Println(time.Date(2000, 8, 1, 1, 1, 1, 1, time.Now().Local().Location()))
	fmt.Println(data, data2)
}
