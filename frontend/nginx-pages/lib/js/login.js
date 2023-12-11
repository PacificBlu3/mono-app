async function submitLogin() {
    const http_req = fetch(window.origin + '/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            email: document.querySelector('input[type="email"]').value,
            password: document.querySelector('input[type="password"]').value
        })
    })   
}