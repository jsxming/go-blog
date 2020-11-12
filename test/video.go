package main

import (
	"net/http"
)


func main() {
	//http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
	//	video, err := os.Open("C:\\Users\\EDZ\\Desktop\\go_proj\\go_blog\\storage\\uploads\\test.mp4")
	//	if err != nil {
	//		log.Printf("Error when try to open file: %v", err)
	//		//sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
	//		return
	//	}
	//
	//	w.Header().Set("Content-Type", "video/mp4")
	//	http.ServeContent(w, r, "", time.Now(), video)
	//
	//	defer video.Close()
	//})

	http.Handle("/", http.FileServer(http.Dir("test")))
	http.ListenAndServe(":9000", nil)

	//http.Handle("/file",http.FileServer(http.Dir("storage")))
	//
	//http.ListenAndServe(":9000",nil)
}