import './Table.css';
import { RequestToS } from '../../Hooks/UseGroupshook';
import { GetCookie } from '../../../Profile/ProfileData';

function Table({ socket, data }) {
    return (
        <>
            <div className='groups-List-header'>
                <h2>Groups List</h2>
            </div>

            <div id='table' className='groups-List-Table'>
                <table>
                    {data.map((ele) => (
                        <tbody
                            className={'groups-List-tbody'}
                            key={ele.GroupID}>
                            <tr
                                onClick={() => {
                                    socket &&
                                        RequestToS(
                                            GetCookie('session_token').split(
                                                '&'
                                            )[0],
                                            ele.Admin,
                                            socket,
                                            'groupInvitationJoin',
                                            ele.GroupID
                                        );
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
