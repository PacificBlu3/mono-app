async function submitLogin() {
    const response = fetch(window.origin + '/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            email: document.querySelector('input[type="email"]').value,
            password: document.querySelector('input[type="password"]').value
        })
    })   
    const status = await response.status;
    if(status == 200) {
        window.location.href = window.location.href
    }
}