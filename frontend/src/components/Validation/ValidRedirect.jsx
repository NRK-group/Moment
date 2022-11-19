import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';

export default function ValidRedirect({children}) {
    const navigate = useNavigate();
    //Query the endpoint and check if the cookie present is valid
    const [logged, setLogged] = useState(true);
    //if cookie is valid show children else redirect
    useEffect(() => {
        fetch('http://localhost:5070/validate', { credentials: 'include' })
            .then((resp) => {
                return resp.text();
            })
            .then((response) => {
                if (response === 'Validated') {
                    navigate('/home')
                    return
                }
                setLogged(false)
                
            });
    }, []);
    return logged ? null : children
}
