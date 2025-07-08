
import React, { useState, useEffect } from 'react'

const API_URL = 'http://localhost:8080/api'

export default function App() {
  const [urls, setUrls] = useState<any[]>([])

  useEffect(() => {
    fetch(\`\${API_URL}/urls\`, {
      headers: {
        'Authorization': 'Bearer secret-token'
      }
    })
      .then(res => res.json())
      .then(data => setUrls(data.urls))
  }, [])

  return (
    <div className="p-4">
      <h1 className="text-2xl mb-4">URL Dashboard</h1>
      <table className="w-full border">
        <thead>
          <tr>
            <th className="border px-2">ID</th>
            <th className="border px-2">Title</th>
          </tr>
        </thead>
        <tbody>
          {urls.map((url: any, idx: number) => (
            <tr key={idx}>
              <td className="border px-2">{url.id}</td>
              <td className="border px-2">{url.title}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}
