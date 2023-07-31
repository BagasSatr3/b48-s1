// Mengirim email tanpa membuka aplikasi 

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

// Mengirim email membuka aplikasi 

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

// function untuk validasi form

function validateform() {
    let name = document.getElementById("name").value;
    let image = document.getElementById("image").value;
    let description = document.getElementById("description").value;
    let startDate = document.getElementById("startDate").value;
    let endDate = document.getElementById("endDate").value;
    let tech = document.querySelectorAll(".check:checked");

    if (name == "") {
        return alert("Please fill project name!");
    }
    if (startDate == "") {
        return alert("Please fill project start date!");
    }
    if (endDate == "") {
        return alert("Please fill project end date!");
    }
    if (description == "") {
        return alert("Please fill project description!");
    }
    if (tech.length === 0) {
        return alert("Please check atleast one technology!");
    }
    if (image == "") {
        return alert("Please fill project image!");
    }
}

let id = 1;

function increment() {
    id++;
    return id;
}

let project = []

// Menambahkan Project

function addProject(event) {
    event.preventDefault();

    let name = document.getElementById("name").value;
    let image = document.getElementById("image").files;
    let description = document.getElementById("description").value;

    const nodejsIcon = '<i class="fa-brands fa-node-js"></i>';
    const reactjsIcon = '<i class="fa-brands fa-react"></i>';
    const vuejsIcon = '<i class="fa-brands fa-vuejs"></i>';
    const javacscriptIcon = '<i class="fa-brands fa-square-js"></i>';

    let cnodejs = document.getElementById("nodejs").checked ? nodejsIcon : "";
    let cnreactjs = document.getElementById("reactjs").checked ? reactjsIcon : "";
    let cvuejs = document.getElementById("vuejs").checked ? vuejsIcon : "";
    let cjavacscript = document.getElementById("javascript").checked ? javacscriptIcon : "";

    let startDate = document.getElementById("startDate").value;
    let endDate = document.getElementById("endDate").value;

    let start = new Date(startDate);
    let end = new Date(endDate);

    if (start > end) {
        alert("Please input date correctly!");
    }

    let diff = (end - start) / 1000;
    let min = Math.abs(Math.round(diff / 60));
    let hour = Math.abs(Math.round(min / 60));
    let day = Math.abs(Math.round(hour / 24));
    let week = Math.abs(Math.round(day / 7));
    let month = Math.abs(Math.round(week / 4));
    let year = Math.abs(Math.round(month / 12));
    let duration = "";

    if (day == 1) {
        duration = day + " day";
    } else if (day < 7) {
        duration = day + " days";
    }

    if (day == 7) {
        duration = week + " week";
    } else if (day < 14 && day > 7) {
        duration = week + " week";
    } else if (day >= 14 && day < 30) {
        duration = week + " weeks";
    }
    
    if (day == 30) {
        duration = month + " month";
    } else if (day < 30 && day > 60) {
        duration = month + " month";
    } else if (day >= 30 && day < 365) {
        duration = month + " months";
    }
    
    if (day == 365) {
        duration = year + " year";
    }else if (day < 730 && day > 365) {
        duration = year + " year";
    }else if (day >= 730) {
        duration = year + " years";
    }

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
        duration,
        id
    }
    
    
    project.push(projects);
    renderProject();
    console.log(project)
    
}

// Menampilkan Project

function renderProject() {
    document.getElementById("project").innerHTML = "";
    for(let i = 0; i < project.length; i++) {
        document.getElementById("project").innerHTML += `
        

        <div class="col col-md-6 col-sm-12 col-lg-6 col-xl-4 mb-4">
            <div class="card shadow" style="width: 22rem;">
                <img src="${project[i].imageURL}" alt="" srcset="" class="rounded m-4">
                <div class="mx-4 mb-4">
                <h4 class="mt-2 text-start"><button class="btn fs-3" type="button" data-bs-toggle="modal" data-bs-target="#${project[i].id}">${project[i].name}</button></h4>
                    <p class="fw-light fs-6">duration : ${project[i].duration}</p>
                    <p class="" maxlength="100" style="text-align: justify;overflow: hidden; white-space: nowrap; text-overflow: ellipsis;">${project[i].description}</p>
                    <div class="fs-4 row">
                        <div class="col-1">
                            ${project[i].cnodejs}
                        </div>
                        <div class="col-1">
                            ${project[i].cvuejs}
                        </div>
                        <div class="col-1">
                            ${project[i].cjavacscript}
                        </div>
                        <div class="col-1">
                            ${project[i].cnreactjs}
                        </div>
                    </div>
                    <div class="row text-center mt-4">
                        <div class="col">
                            <button class="btn btn-dark rounded-pill px-5">Edit</button>
                        </div>
                        <div class="col">
                            <button class="btn btn-dark rounded-pill px-5">Delete</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>


<div class="modal fade" id="${project[i].id}" tabindex="-1" aria-labelledby="${project[i].name}Label" aria-hidden="true">
  <div class="modal-dialog modal-xl">
    <div class="">
    <div class="container" id="detail">
    <div class=" card">
        <h2 class="d-flex justify-content-center m-4 fs-2">${project[i].name}</h2>
        <div class="row m-4">
            <div class="col-8  border-0">
                <img src="${project[i].imageURL}" alt="" srcset="" class="card-img-top detail-img">
            </div>
            <div class="col-4 mt-4">
                <h4 class="">Duration</h4>
                <p><i class="fa-solid fa-calendar-days"></i> ${project[i].startDate} - ${project[i].endDate}</p>
                <p><i class="fa-sharp fa-regular fa-clock"></i>  ${project[i].duration}</p>
                <h4 class="">Technologies</h4>
                        ${project[i].cnodejs}
                        ${project[i].cvuejs}
                        ${project[i].cjavacscript}
                        ${project[i].cnreactjs}
            </div>
        </div>
       <div class="m-4">
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

// Instant copy description

function copyClipboard() {
    // Get the text field
    var copyText = document.getElementById("myInput");
  
    // Select the text field
    copyText.select();
    copyText.setSelectionRange(0, 99999); // For mobile devices
  
     // Copy the text inside the text field
    navigator.clipboard.writeText(copyText.value);
  
    // Alert the copied text
    alert("Copied the text: " + copyText.value);
  }
