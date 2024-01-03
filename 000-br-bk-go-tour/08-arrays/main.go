package main

import "fmt"

func main() {
	var animals [5]string
	animals[0] = "Aardvark"
	animals[1] = "Platypus"
	animals[2] = "Narwhal"
	animals[3] = "Tarsier"
	animals[4] = "Kakapo"

	fmt.Println("------- CONTIGUOUS MEMORY")
	// an array creates a contiguous block of memory
	for i, v := range animals {
		fmt.Printf("%d \t %v \t %p \n", i, v, &animals[i])
	}

	fmt.Println("------- CONTIGUOUS MEMORY")
	for i, v := range animals {
		fmt.Printf("%d \t %v \t %p \t %p \n", i, v, &v, &animals[i])
	}

	fmt.Println("--------- BIRDS")

	birds := [5]string{"'I'iwi", "Nēnē", "'Apapane", "'Elepaio", "Pueo"}
	fmt.Println(birds)

	fmt.Println("--------- RANGE LOOP ---------")
	// COPIES THE VALUE from the `birds` array to the variable `v`
	// `v` gets its own copy of the value
	for i, v := range birds {
		fmt.Println(i, v)
	}

	fmt.Println("--------- FOR LOOP ---------")
	// access values BY INDEX POSITION
	for i := 0; i < len(birds); i++ {
		fmt.Println(i, birds[i])
	}

	// changing v doesn't change `birds`
	fmt.Println("--------- v changed ---------")
	for _, v := range birds {
		v = "v changed"
		fmt.Println(v)
	}
	fmt.Println("birds", birds)

	// changing the value at an index position does change `birds`
	fmt.Println("--------- birds changed ---------")
	for i := range birds {
		birds[i] = "birds changed"
	}
	fmt.Println("birds", birds)


	psychologists := [5]string{"Alpert", "C.Jung", "Piaget", "Rogers", "Skinner"}

	for i, v := range psychologists {
		psychologists[1] = "MASLOW"
		fmt.Printf("%d \t %s \t \t \t %s \n", i, v, psychologists[1])
	}
	fmt.Println(psychologists)

	psychologists = [5]string{"Alpert", "C.Jung", "Piaget", "Rogers", "Skinner"}
	for i, v := range &psychologists {
		psychologists[1] = "MASLOW"
		fmt.Printf("%d \t %s \t \t \t %s \n", i, v, psychologists[1])
	}
	fmt.Println(psychologists)
}
