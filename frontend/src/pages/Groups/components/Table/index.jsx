import './table.css';
// import {RequestToJoin} from '../../hooks/useGroupshook'

function Table({ data }) {
    
    return (
        <>
            <div className='groups-List-header'>
                <h2>Groups List</h2>
            </div>

            <div id='table' className='groups-List-Table'>
                <table>
                    {data.map((ele) => (
                        <tbody className={ele.member?'groups-List-Block': 'groups-List-tbody'} key={ele.GroupID}>
                            <tr onClick={() => console.log(ele)}>
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
