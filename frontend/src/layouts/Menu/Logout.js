export default function Logout(redir, setauth) {
   let logged =  fetch('http://localhost:5070/logout', {
        credentials: 'include'}).then( async response => {
            return  await response.text()     
        }). then( resp => {
            if (resp === "Logged Out") {
                setauth(false)
                redir('/')
                return true
            }
            return false
        })
}