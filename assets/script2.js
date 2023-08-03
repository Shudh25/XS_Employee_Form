
const formE1 = document.getElementById('employee_form')

formE1.addEventListener('submit', (event) => {
    event.preventDefault();

    const data = new FormData(formE1);
    data.append("Resume", resume.files[0]);

    fetch('/sendData', {
        method: 'POST',
        // headers: {
        //     'Content-Type': 'multipart/form-data'
        //     // 'Content-Type': 'application/x-www-form-urlencoded'
        // },
        body: data
    }).then(res => res.json())
        .then(data => console.log(data))
        .then(error => console.log(error))

});
