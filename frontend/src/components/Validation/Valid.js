
export default function CheckCookie(setter) {
    //Send a get request to check if the user is valid----let result = await 
    fetch("http://localhost:5070/validate", {credentials: 'include'}).then( resp => {
        return resp.text()
    }).then(response => {
        console.log(response)
        if (response === 'Validated'){
            setter(true)
            return
        }
        setter(false)
        return
    })
    // console.log(result)
    // return result
}