function generateResponse() {
    const userInput = document.getElementById("userInput").value;
    alert("AI is thinking. It takes a while")

    fetch(`http://localhost:3030/generate/${encodeURIComponent(userInput)}`)
        .then(response => response.text())
        .then(data => {
            document.getElementById("responseContainer").innerText = data;

            const linkContainer = document.createElement("div");
            const link = document.createElement("a");
            link.href = `/response/${data}`;
            link.innerText = "Go to Response";
            linkContainer.appendChild(link);
            document.body.appendChild(linkContainer);
        })
        .catch(error => {
            console.error('Error generating response:', error);
        });
}