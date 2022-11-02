import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
export default function Validation({setAuth, children}) {
    const navigate = useNavigate();
    //Query the endpoint and check if the cookie present is valid
    const [valid, setValid] = useState(false);
    //if cookie is valid show children else redirect
    useEffect(() => {
        fetch('http://localhost:5070/validate', { credentials: 'include' })
            .then((resp) => {
                return resp.text();
            })
            .then((response) => {
                if (response === 'Validated') {
                    setValid(true);
                    return;
                }
                navigate('/')
            });
    }, []);
    return valid ? children : null ;
}
