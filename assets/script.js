const formE1 = document.getElementById('employee_form')
const phone = document.getElementById('phone')
const resume = document.getElementById('resume')
const email = document.getElementById('email')
const errorElement = document.getElementById('error')

formE1.addEventListener('submit', (e) => {


    function formValidation() {

        let error_messages = []

        // PHONE VALIDATIONS

        var numbers = /^[0-9]+$/;

        if (phone.value.match(numbers) == null) {
            error_messages.push("Enter Number only in Phone")
            console.log("hello")
        }

        // If the Entered Phone Is less than 10 Digit
        let phone_length = phone.value.toString().length
        if (phone_length < 10) {
            error_messages.push("PLease Enter valid Phone Number")
        }


        //FILE VALIDATIONS

        var InputFile = resume;
        var filePath = InputFile.value;

        // Allowing file type
        var allowedExtensions = /(\.pdf|\.png)$/i;

        if (!allowedExtensions.exec(filePath)) {
            error_messages.push('Invalid file type');
            InputFile.value = '';
        }


        //EMAIL VALIDATIONS

        var validRegex = /^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@(xenonstack.in|xenonstack.com)$/;
        var receivedEmail = email.value.trim();

        //Checking Custom Domain Vaidation
        if (!(receivedEmail.toLowerCase().match(validRegex))) {
            error_messages.push("Please enter valid email address.")
        }


        // Printing All error_messages
        if (error_messages.length > 0) {
            e.preventDefault()
            errorElement.innerText = error_messages.join(', ')
            errorElement.style.height = "10px";
            console.log(error_messages.length)
        } else {
            return true
        }
    }


    if (formValidation()) {
        console.log(error_messages.length)
        console.log(error_messages.length <= 0)

        e.preventDefault();

        const data = new FormData(formE1);
        data.append("Resume", resume.files[0]);

        fetch('/sendData', {
            method: 'POST',
            body: data
        }).then(res => res.json())
            .then(data => console.log(data))
            .then(error => console.log(error))
    }


})

