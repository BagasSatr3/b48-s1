function sendMail(){
    var params = {
        name : document.getElementById("name").value,
        email : document.getElementById("email").value,
        number : document.getElementById("number").value,
        subject : document.getElementById("subject").value,
        message : document.getElementById("message").value
    };
    
    const serviceID = "service_l62cm3s";
    const templateID = "template_9adh69l";
        
    emailjs
    .send(serviceID, templateID, params)
    .then((res) =>{
        document.getElementById("name").value = "";
        document.getElementById("email").value = "";
        document.getElementById("number").value = "";
        document.getElementById("subject").value = "";
        document.getElementById("message").value = "";
        console.log(res);
        alert("your message sent successfully");
    })
    .catch((err) => console.log(err));
}

function mail() {
    let name = document.getElementById("name").value
    let email = document.getElementById("email").value
    let number = document.getElementById("number").value
    let subject = document.getElementById("subject").value
    let message = document.getElementById("message").value

    if(name == "") {
        return alert("Please insert your Name")
    } else if(email == "") {
        return alert("Please insert your Email address")
    } else if(number == "") {
        return alert("Insert your phone number, please")
    } else if(subject == "") {
        return alert("Select your subject, please")
    } else if(message == ""){
        return alert("Typing your message, please")
    }

    const destination = "bagassatjin23@gmail.com";
    let a = document.createElement("a")
    a.setAttribute('href', `mailto:${destination}?subject=${subject}&body= Hello, my name is ${name} , my reason contact you is ${message}, contact me at ${number}`)
    a.click()
    

    let data = {
        name,
        email,
        number,
        subject,
        message
    }

    console.log(data)
}