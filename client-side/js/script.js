const submitButton = document.getElementById("submit")
submitButton.addEventListener("click", function () {

    // Entradas
    const cpfInput  = document.getElementById("cpf")
    const nameInput = document.getElementById("name")
    const ageInput  = document.getElementById("age")

    // Saída
    const askTime = document.getElementById("ask-time")

    let data = {
        Cpf:  cpfInput.value,
        Name: nameInput.value,
        Age:  ageInput.value,
    };

    console.log(data)

    let eleentType = document.getElementsByTagName("action")

    switch (eleentType) {
        case : 
    }

    !(async () => {
        const rawResponse = await fetch('/tim', {
            method: "POST",
            headers: {"Content-type": "application/json;charset=UTF-8",
                'Accept': 'application/json',
            //     'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)

        }).then((response) => {
            response.text().then(function (data) {
                let result = JSON.parse(JSON.stringify(data));
                console.log(result)
            });
        }).catch((error) => {
            console.log(error)
        });

    })(data);
})

