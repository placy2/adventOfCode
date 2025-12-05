document
  .getElementById("runnerForm")
  .addEventListener("submit", function (event) {
    event.preventDefault();
    const form = this;

    fetch(form.action, {
      method: "POST",
      body: new FormData(this),
    })
      .then((response) => response.json().then(data => ({ status: response.status, data })))
      .then(({ status, data }) => {
        if (status === 200) {
          document.getElementById("result").innerText = data.message;
        } else {
          document.getElementById("result").innerText = data.error || "An error occurred";
        }
      })
      .catch((error) => {
        document.getElementById("result").innerText = "Network error: " + error.message;
        console.error("Error:", error);
      });
  });