export default function ApplicationsTable({ applications }) {
  return (
    <table>
      <Header />
      <Body applications={applications} />
    </table>
  )
}

function Header() {
  return (
    <thead>
      <tr>
        <th>Date</th>
        <th>Company</th>
        <th>Role</th>
        <th>Location</th>
        <th>Status</th>
        <th></th>
      </tr>
    </thead>
  );
}

function Body({ applications }) {
  return (
    <tbody>
      {applications.map(app => (
        <tr key={app.Id} className="application">
          <td>{app.Date}</td>
          <td>{app.Company}</td>
          <td>{app.Role}</td>
          <td>{app.Location}</td>
          <td>{app.Status}</td>
          <td>
            <button onClick={() => { console.log(`Edit button pressed on application with id of ${app.Id}`) }}>Edit</button>
          </td>
        </tr>
      ))}
    </tbody>
  )
}
