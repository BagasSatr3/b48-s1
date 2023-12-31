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
        // console.log(response)
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
        <div class="col col-md-6 col-sm-12 col-lg-6 col-xl-4 mb-4">
            <div class="mt-2 card shadow-sm" style="width: 21em;">
                <img class="p-3 rounded-5 " style="object-fit: cover; width: 21em; height: 16rem;" src="${card.image}" alt="">
                <div class="mt-2 mx-4">
                    <p class="fst-italic fw-bold">"${card.quote}"</p>
                    <br>
                    <div class="row ">
                        <p class="col text-start">${card.rating} <i class="fa-solid fa-star"></i></p>
                        <p class="col text-end fw-bold text-wrap">-${card.user}</p>
                    </div>
                </div>
            </div>
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