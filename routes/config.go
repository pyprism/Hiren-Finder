package routes


func Init() {
	r := NewRouter()
	err := r.Run("127.0.0.1:4200")
	if err != nil {
		panic(err)
	}
}
