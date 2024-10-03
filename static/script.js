function addImageUrlField() {
    const container = document.getElementById("image")
    const input = document.createElement("input")
    input.type = "url"
    input.name = "image"
    input.placeholder = "Enter image URL"
    container.appendChild(input)
    const deleteButton = document.createElement("button")
    deleteButton.type = "button"
    deleteButton.textContent = "Remove"
    deleteButton.onclick = function() {
        container.removeChild(input)
        container.removeChild(deleteButton)
    }
    container.appendChild(deleteButton)
}