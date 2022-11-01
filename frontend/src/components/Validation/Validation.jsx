import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import CheckCookie from './Valid';
export default function Validation(props) {
    const navigate = useNavigate();
    //Query the endpoint and check if the cookie present is valid
    const [valid, setValid] = useState(true);
    console.log('YOUR ARE CALLING VALIDATION');
    //if cookie is valid show children else redirect
    // setValid();
    // CheckCookie(setValid);
    // console.log({valid})
    // if (!valid) {
    //     navigate('/');
    //     return;
    // }
    // console.log(props.children, valid.then(resp => resp).then(response => response));

    return props.children;
}
