$("video").each((i, e) => {
	$(e).hover(startVideo, stopVideo)
})
function startVideo(e){
	e.target.load()
	e.target.play()
}
function stopVideo(e){
	e.target.pause()
}
