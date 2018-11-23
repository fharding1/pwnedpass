// Package pwnedpass provides a library for accessing the Pwned Password API.
//
// Usage:
// 	count, err := pwnedpass.DefaultClient.Count("password")
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	fmt.Println(count)
//
// You can also create your own Client for use with a different API or to use
// a different HTTP client:
//
//  client := pwnedpass.ClientV2{
// 	 HTTPClient: &http.Client{
// 		 Timeout: time.Second * 2,
// 	 },
//  }
//
//  count, err := client.Count("password")
//  if err != nil {
// 	 panic(err)
//  }
//
// fmt.Println(count)
package pwnedpass
