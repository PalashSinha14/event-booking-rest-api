// LOGIN
const loginForm = document.getElementById("loginForm");

if(loginForm){
loginForm.addEventListener("submit", async function(e){

e.preventDefault();

const email = document.getElementById("email").value;
const password = document.getElementById("password").value;

const res = await fetch("/login",{
method:"POST",
headers:{
"Content-Type":"application/json"
},
body:JSON.stringify({
email:email,
password:password
})
});

const data = await res.json();

if (res.ok) {
    window.location.href = "/dashboard";
} else {
    alert(data.message);
}

//alert(JSON.stringify(data));

});
}


// SIGNUP
const signupForm = document.getElementById("signupForm");

if(signupForm){

signupForm.addEventListener("submit", async function(e){

e.preventDefault();

const name = document.getElementById("name").value;
const email = document.getElementById("email").value;
const password = document.getElementById("password").value;

const res = await fetch("/signup",{
method:"POST",
headers:{
"Content-Type":"application/json"
},
body:JSON.stringify({
name:name,
email:email,
password:password
})
});

const data = await res.json();

alert(JSON.stringify(data));

});
}