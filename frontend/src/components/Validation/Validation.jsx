import { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
export default function Validation(props) {
    const navigate = useNavigate();
    //Query the endpoint and check if the cookie present is valid
    let valid = CheckCookie();
    console.log('YOUR ARE CALLING VALIDATION');
    //if cookie is valid show children else redirect
    useEffect(() => {
        if (!valid) navigate('/');
    }, [valid]);
    return valid ?? props.children;
}
