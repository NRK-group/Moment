import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
export default async function Validation(auth) {
    // const [authorised, setAuthorised] = useState(false);
    // const navigate = useNavigate();
    // useEffect(() => {
        return await fetch('http://localhost:5070/validate', {
            credentials: 'include',
        }).then(async (resp) => {
            const response = await resp.text();
            // if (response !== 'Validated') {
            //     setAuthorised(false);
            //     navigate('/');
            //     return;
            // }
            // setAuthorised(true);
            // return;
        });
    // }, [auth]);
    // return authorised;
}
