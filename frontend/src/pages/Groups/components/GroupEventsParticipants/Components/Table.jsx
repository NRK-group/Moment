import './tableParticipants.css';

function TabletableParticipants({ data }) {
   
    return (
        <>
            <div className='groups-List-header'>
                <h2>Group Events Participants</h2>
            </div>

            <div id='table' className='groups-List-Table'>
                <table>
                    {data && data.map((ele) => (
                        <tbody key={ele.UserId}>
                            <tr>
                                <td>{ele.Name}</td>
                            </tr>
                        </tbody>
                    ))}
                </table>
            </div>
        </>
    );
}

export default TabletableParticipants;
