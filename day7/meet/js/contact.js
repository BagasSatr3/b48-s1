function getData(event) {
    event.preventDefault();

    let name = document.getElementById("name").value
    let email = document.getElementById("email").value
    let number = document.getElementById("number").value
    let subject = document.getElementById("subject").value
    let message = document.getElementById("message").value

    let objectData = {
        name,
        email,
        number,
        subject,
        message
    }

    let arrayData = [name, email, number, subject, message]

    console.log(objectData)

    if (name === "") {
        return alert('Name harus diisi!')
    } else if (email === "") {
        return alert('Email harus diisi!')
    } else if (number === "") {
        return alert('Phone harus diisi!')
    } else if (subject === "") {
        return alert('Subject harus diisi!')
    } else if (message === "") {
        return alert('Message harus diisi!')
    }

    const emailReceiver = "bagassatjin23@gmail.com"

    let a = document.createElement('a')
    a.href = `mailto:${emailReceiver}?subject=${subject}&body=Halo nama saya ${name},\n${message}, silahkan kontak saya di nomor berikut : ${number}`
    a.click()
    
}