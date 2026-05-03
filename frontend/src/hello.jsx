import { useEffect, useState } from "react";

export default function Hello() {
  const [text, setText] = useState("");
  useEffect(() => {
    async function getApplications() {
      try {
        const response = await fetch('api/v1/applications', { method: 'GET' });
        
        if (!response.ok) {
          setText("");
          return;
        }
        
        setText(await response.text());
      } catch (e) {
        setText("");
        console.error(e);
      }
    }
    getApplications();
  }, []);

  return (
    <p>{text}</p>
  );
}