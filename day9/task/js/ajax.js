const promise = new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest()

    xhr.open("GET", "https://api.npoint.io/87896768813bcf08ce4f", true)
    xhr.onload = function () {
        // HTTP status code
    if (xhr.status === 200) {
        resolve(JSON.parse(xhr.responseText))
    } else if (xhr.status >= 400) {
        reject("Error loading data")
    }
    }
    xhr.onerror = function () {
        reject("Network error")
    }
    xhr.send()
})

let testimonialData = []

async function getData() {
    try {
        const response = await promise
        console.log(response)
        testimonialData = response
        allTestimonial()
    } catch (err) {
        console.log(err)
    }
}

getData()

function allTestimonial() {
    let testimonialHTML = ""

    testimonialData.forEach((card) => {
        testimonialHTML += `
        <div class="cards">
            <img src="${card.image}" alt="">
            <p class="quote">"${card.quote}"</p>
            <br>
            <br>
            <p class="user">- ${card.user}</p>
            <p class="star">${card.rating} <i class="fa-solid fa-star"></i></p>
        </div>
        `
    })

    document.getElementById("testimonials").innerHTML = testimonialHTML
}

function filterTestimonial(rating) {
    let filteredTestimonialHTML = ""

    const filterdData = testimonialData.filter((card) => {
       return card.rating === rating
    })

    filterdData.forEach((card) => {
        filteredTestimonialHTML += `
        <div class="cards">
            <img src="${card.image}" alt="">
            <p class="quote">"${card.quote}"</p>
            <br>
            <br>
            <p class="user">- ${card.user}</p>
            <p class="star">${card.rating} <i class="fa-solid fa-star"></i></p>
        </div>
        `
    })

    document.getElementById("testimonials").innerHTML = filteredTestimonialHTML
}