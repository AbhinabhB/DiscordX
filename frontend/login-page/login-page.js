const loginForm = document.getElementById("login-form");
const loginButton = document.getElementById("login-form-submit");
const loginErrorMsg = document.getElementById("login-error-msg");

function redirect_sign_in() {
    location.assign("../sign-in-page/sign-in-page.html");
}

loginButton.addEventListener("click", (e) => {
    e.preventDefault();
    const email = loginForm.email.value;
    const password = loginForm.password.value;

    if(email == "test_email" && password == "test_password") {
        alert("You have successfully logged in.");
        location.reload();
    } else {
        alert("Wrong Credentials!");
    }
})