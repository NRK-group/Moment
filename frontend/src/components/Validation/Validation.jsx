import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import CheckCookie from './Valid';
export default function Validation(props) {
    const navigate = useNavigate();
    //Query the endpoint and check if the cookie present is valid
    const [valid, setValid] = useState(true);
    console.log('YOUR ARE CALLING VALIDATION', CheckCookie());
    //if cookie is valid show children else redirect
    useEffect(() => {
        setValid(CheckCookie());
        if (!valid) {
            navigate('/');
        }
    }, [valid]);

    return props.children;
}
