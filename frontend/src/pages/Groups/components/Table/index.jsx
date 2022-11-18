import './table.css';

function Table({ data }) {
    console.log({ data });
    return (
        <>
        <div className='groups-List-header'>
        <h2>Groups List</h2>
        </div>
            
            <div id='table' className='groups-List-Table'>
                <table>
                    {data.map((ele) => (
                        <tbody key={ele.GroupID}>
                            <tr onClick={()=>console.log(ele)}>
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
