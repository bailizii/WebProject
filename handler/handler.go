package handler

import (
	"WebProject/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var people []model.Person = []model.Person{}

func AddData() {
	people = append(people, model.Person{ID: "1", Name: "张三", Phone:"18301236711"})
	people = append(people, model.Person{ID: "2", Name: "李四", Phone:"16201236711"})
	people = append(people, model.Person{ID: "3", Name: "王五", Phone:"18301236711"})

}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	// get the ID of the person from the route parameter
	var idParam string = mux.Vars(r)["id"] //用键“ID”从该映射中检索值; 在Go中的地图中建立索引可以有选择地返回两个值，（1）一个映射到提供的键的值，（2）一个布尔值，指示一个值是否实际映射到键。
	id, err := strconv.Atoi(idParam)

	if err != nil {
		// there was an error
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	// error checking
	if id >= len(people) {
		w.WriteHeader(404)
		w.Write([]byte("Nothing found with specified ID"))
		return
	}

	person := people[id]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func GetAllPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	// get Item value from the JSON body
	var newPerson model.Person
	json.NewDecoder(r.Body).Decode(&newPerson)

	//body,_:= ioutil.ReadAll(r.Body)  // =====》 你把body直接读出来，看是什么  因为int那个bug在差错的时候做的

	var text string = newPerson.ID
	//fmt.Println(len(people))
	for i := 0; i< len(people); i++{
		if people[i].ID == text {
			return
		}else {
			continue
		}
	}

	people = append(people, newPerson)
	//fmt.Println(people)

	//fmt.Println(string(body))
	//fmt.Println(newPerson)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(people)
}



func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	// get the ID of the person from the route parameters
	var idParam string = mux.Vars(r)["id"]

	id, err := strconv.Atoi(idParam) //convert string to integer
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	// error checking
	if id >= len(people) {
		w.WriteHeader(404)
		w.Write([]byte("Nothing found with specified ID"))
		return
	}

	// get the value from JSON body
	var updatedPerson model.Person
	//fmt.Println("....")
	//body,_:= ioutil.ReadAll(r.Body)  // =====》 你把body直接读出来，看是什么  因为int那个bug在差错的时候做的
	//fmt.Println(string(body))

	json.NewDecoder(r.Body).Decode(&updatedPerson)

	people[id] = updatedPerson

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedPerson)
}


func DeletePerson(w http.ResponseWriter, r *http.Request) {
	// get the ID of the person from the route parameters
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	// error checking
	if id >= len(people) {
		w.WriteHeader(404)
		w.Write([]byte("Nothing found with specified ID"))
		return
	}

	// Delete the person from the slice
	// https://github.com/golang/go/wiki/SliceTricks#delete
	people = append(people[:id], people[id+1:]...)

	w.WriteHeader(200)
}