import LogoutComp from "./LogoutComp";

export default function Logout(redir, setauth) {
    fetch('http://localhost:5070/logout', {
        credentials: 'include',
    })
        .then(async (response) => {
            return await response.text();
        })
        .then((resp) => {
            if (resp === 'Logged Out') {
                location.reload()
                setauth(false)
                redir('/')
                return true;
            }
            return false;
        });
}
