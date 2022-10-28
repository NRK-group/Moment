export default function Validation(props) {
    //Query the endpoint and check if the cookie present is valid
    let valid = false
    //if cookie is valid show children else redirect
  return valid 
  ? props.children
  : <Navigate to="/login" replace />
}
