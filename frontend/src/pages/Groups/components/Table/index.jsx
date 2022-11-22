import './table.css';
import {RequestToS} from '../../hooks/useGroupshook'
import { GetCookie } from '../../../profile/ProfileData'; 

function Table({ socket, data }) {

   
    
    return (
        <>
            <div className='groups-List-header'>
                <h2>Groups List</h2>
            </div>

            <div id='table' className='groups-List-Table'>
                <table>
                    {data.map((ele) => (
                        <tbody className={ele.member?'groups-List-Block': 'groups-List-tbody'} key={ele.GroupID}>
                            <tr onClick={() =>{
                                 RequestToS(GetCookie('session_token').split('&')[0], ele.Admin, socket, "groupInvitationJoin", ele.GroupID)
                                console.log(ele)
                            }}>
                                <td>{ele.Name}</td>
                            </tr>
                        </tbody>
                    ))}
                </table>
            </div>
        </>
    );
}

export default Table;
