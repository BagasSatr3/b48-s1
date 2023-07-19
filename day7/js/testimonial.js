class Testimonial {
    #quote = ""
    #image = ""

    constructor(quote, image) {
        this.#quote = quote
        this.#image = image
    }

    get quote() {
        return this.#quote
    }

    get image() {
        return this.#image
    }

    get user() {
        throw new Error('there is must be user to make testimonials')
    }

    get testimonialHTML() {
        return `
        <div class="cards">
            <img src="${this.image}" alt="">
            <p class="quote">"${this.quote}"</p>
            <br>
            <br>
            <p class="user">- ${this.user}</p>
        </div>
        `
    }
}

class UserTestimonial extends Testimonial {
    #user = ""

    constructor(user, quote, image) {
        super(quote, image)
        this.#user = user
    }

    get user() {
        return this.#user + `<i class="fa-solid fa-user" style="margin:4px; margin-left:10px"></i>`
    }
}

class CompanyTestimonial extends Testimonial {
    #company = ""

    constructor(company, quote, image) {
        super(quote, image)
        this.#company = company
    }

    get user() {
        return this.#company + `<i class="fa-sharp fa-regular fa-building" style="margin:4px; margin-left:10px"></i>`
    }
}

const testimonial1 = new UserTestimonial("Philia", "What is this thing?", "https://w.wallha.com/ws/14/9qHfZhkC.png")

const testimonial2 = new UserTestimonial("Catto", "Gimme more treats human!", "https://www.thesprucepets.com/thmb/APYdMl_MTqwODmH4dDqaY5q0UoE=/750x0/filters:no_upscale():max_bytes(150000):strip_icc():format(webp)/all-about-tabby-cats-552489-hero-a23a9118af8c477b914a0a1570d4f787.jpg")

const testimonial3 = new CompanyTestimonial("Tuxedo Catto", "Look at you, staring at your laptop all day. Go touch some grass.", "https://i2.wp.com/felineliving.net/wp-content/uploads/2018/12/Tuxedo-Cats.jpg")

let testimonialData = [testimonial1, testimonial2, testimonial3]

let testimonialHTML = ""

for (let i = 0; i < testimonialData.length; i++) {
    testimonialHTML += testimonialData[i].testimonialHTML
}

document.getElementById("testimonials").innerHTML = testimonialHTML