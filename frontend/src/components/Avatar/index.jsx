import { useState, useEffect } from 'react';
import './Avatar.css';

function Avatar({ avatarSrc, avatarAlt, styleName, name , avatarRedirect, avatarIcon }) {
    const [src, setSrc] = useState(avatarSrc);
    const [alt, setAlt] = useState(avatarAlt);
    const [nameS, setNameS] = useState(name);
    const [avatarIconS, setAvatarIconS] = useState(name);
    const [styleNameS, setStyleNameS] = useState(styleName);

    useEffect(() => {
        setSrc(avatarSrc)
        setAlt(avatarAlt)
        setStyleNameS(styleName)
        setNameS(name)
        setAvatarIconS(avatarIcon)
    }, [avatarSrc, avatarAlt, styleName, avatarIcon]);


    return (
        <div className='AvatarC'>
            <img className={styleNameS} src={src} alt={alt} onClick={()=>{avatarRedirect}} />
          { avatarIconS && (<span className='AvatarIcon'><i className="fa-solid fa-circle-plus"></i></span>)  }
           { nameS && <p>{nameS}</p> }
            
        </div>
    );
}

export default Avatar;
