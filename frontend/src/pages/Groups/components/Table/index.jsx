import './table.css';

function Table({ data }) {
    console.log({data})
    return (
        <div id='table'>
            <h2>Groups List</h2>
            <table>
                <tbody>
                    <tr>
                        <th>Groups</th>
                    </tr>
                </tbody>
                {data.map((ele) => (
                    <tbody key={ele.GroupID}>
                        <tr>
                            <td>{ele.Name}</td>
                        </tr>
                    </tbody>
                ))}
            </table>
        </div>
    );
}

export default Table;
