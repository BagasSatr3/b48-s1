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

let id = 1

function increment() {
    id++
}

let project = []

function addProject(event) {
    event.preventDefault()

    let name = document.getElementById("name").value
    let image = document.getElementById("image").files
    let description = document.getElementById("description").value

    const nodejsIcon = '<i class="fa-brands fa-node-js"></i>';
    const reactjsIcon = '<i class="fa-brands fa-react"></i>';
    const vuejsIcon = '<i class="fa-brands fa-vuejs"></i>';
    const javacscriptIcon = '<i class="fa-brands fa-square-js"></i>';

    let cnodejs = document.getElementById("nodejs").checked ? nodejsIcon : "";
    let cnreactjs = document.getElementById("reactjs").checked ? reactjsIcon : "";
    let cvuejs = document.getElementById("vuejs").checked ? vuejsIcon : "";
    let cjavacscript = document.getElementById("javascript").checked ? javacscriptIcon : "";

    let startDate = document.getElementById("startDate").value
    let endDate = document.getElementById("endDate").value

    const imageURL = URL.createObjectURL(image[0])
    
    let projects = {
        name,
        description,
        imageURL,
        startDate,
        endDate,
        cnodejs,
        cnreactjs,
        cvuejs,
        cjavacscript,
        id
    }
    
    
    project.push(projects);
    renderProject();
    // console.log(projects)
    
}


function renderProject() {
    document.getElementById("project").innerHTML = "";
    for(let i = 0; i < project.length; i++) {
        document.getElementById("project").innerHTML += `
        <div class="profile col-3 card">
                <div class="mb-3">
                    <img src="${project[i].imageURL}" alt="" srcset="" class="card-img-top">
                    <div class="profile-text">
                    <h4 class="mt-2"><button class="modal-button" type="button" data-bs-toggle="modal" data-bs-target="#${project[i].id}">${project[i].name}</button></h4>
                        <span>duration : 3 month</span>
                        <p class="textw" >${project[i].description}</p>
                        <p>${project[i].startDate} - ${project[i].endDate}</p>
                        ${project[i].cnodejs}
                        ${project[i].cvuejs}
                        ${project[i].cjavacscript}
                        ${project[i].cnreactjs}
                        <div class="row mt-2">
                            <div class="col">
                                <button class="abutton" >Edit</button>
                            </div>
                            <div class="col">
                                <button class="abutton">Delete</button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

<div class="modal fade" id="${project[i].id}" tabindex="-1" aria-labelledby="${project[i].name}Label" aria-hidden="true">
  <div class="modal-dialog modal-xl">
    <div class="">
    <div class="container" id="detail">
    <div class="content-detail card">
        <h2 class="d-flex justify-content-center detail-header">${project[i].name}</h2>
        <div class="row">
            <div class="col-8  border-0">
                <img src="${project[i].imageURL}" alt="" srcset="" class="card-img-top detail-img">
            </div>
            <div class="col-4 mt-4">
                
                <h4 class="">Duration</h4>
                <p><i class="fa-solid fa-calendar-days"></i> ${project[i].startDate} - ${project[i].endDate}</p>
                <p><i class="fa-sharp fa-regular fa-clock"></i>  1 month</p>
                <h4 class="">Technologies</h4>
                        ${project[i].cnodejs}
                        ${project[i].cvuejs}
                        ${project[i].cjavacscript}
                        ${project[i].cnreactjs}
            </div>
        </div>
       <div class="detail-desc">
        <h3 class="mt-4">Description</h3>
        <p class="mt-1" >${project[i].description}</p>
       </div>
    </div>
  </div>
    </div>
  </div>
</div>
        `;
    }
}

// var el_up = document.getElementById("GFG_UP");
// var el_down = document.getElementById("GFG_DOWN");
//     el_up.innerHTML = "Click on button to get ID";
  
// function GFG_click(clicked) {
//     el_down.innerHTML = "ID = "+clicked;
//     el_down.innerHTML = `"ID = "`+clicked;
// }  


