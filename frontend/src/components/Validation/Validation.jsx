import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import CheckCookie from './Valid';
export default function Validation(props) {
    const navigate = useNavigate();
    //Query the endpoint and check if the cookie present is valid
    const [valid, setValid] = useState(false);
    console.log('YOUR ARE CALLING VALIDATION');
    //if cookie is valid show children else redirect
    // setValid();
    // CheckCookie(setValid);
    useEffect(() => {
        fetch('http://localhost:5070/validate', { credentials: 'include' })
            .then((resp) => {
                return resp.text();
            })
            .then((response) => {
                console.log(response);
                if (response === 'Validated') {
                    // setValid(true)
                    setValid(true);
                    return;
                }
                navigate('/')
            });
    });
    // console.log({valid})
    // useEffect(() => {
    // if (!valid) {
    //     return;
    // }
    // });
    // console.log(props.children, valid.then(resp => resp).then(response => response));
    console.log("VALID === ", valid)
    return valid ? props.children : null ;
}
