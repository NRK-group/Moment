import LogoutComp from "./LogoutComp";

export default function Logout(redir, setauth) {
    fetch('http://localhost:5070/logout', {
        credentials: 'include',
    })
        .then(async (response) => {
            return await response.text();
        })
        .then((resp) => {
            console.log({ authorised });
            if (resp === 'Logged Out') {
                location.reload()
                setauth(false)
                setTimeout(redir('/'), 5000)
                return true;
            }
            return false;
        });
}
