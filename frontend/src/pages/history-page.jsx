import { useEffect, useState } from "react";
import { Bar, BarChart, Tooltip, XAxis, YAxis } from "recharts";
import * as Request from "../request.js";

export default function HistoryPage() {
    const [year, setYear] = useState(new Date().getUTCFullYear());
    const [data, setData] = useState(undefined);

    useEffect(() => {
        let ignore = false;
        
        const getData = async () => {
            
            const response = await Request.getData(`${Request.apiUrl}/history/${year}`);
            return await response.json();
        };
        
        getData().then((result) => {
            if (!ignore) setData(result);
        });

        return () => {
            ignore = true;
        }
    }, [year]);

    const handleYearChange = (event) => {
        let value = parseInt(event.target.value, 10);
        if (!Number.isNaN(value)) {
            if (value < 0) value = 0;
            else if (value > 9999) value = 9999;

            setYear(value);
        }
    }

    return (
        <div>
            <h1 className="history-heading">{year} Applications per Month</h1>

            <BarChart
                className="history-chart"
                responsive
                data={data}
                style={{ width: "100%", maxWidth: "900px", aspectRatio: 1.618 }}
            >
                <XAxis dataKey="Month"/>
                <YAxis
                    label={{ value: "Applications", position: "insideLeft", angle: -90, fontWeight: "bold"}}
                    allowDecimals={false}
                />
                <Tooltip />
                <Bar dataKey="Count" fill="dodgerblue"/>
            </BarChart>

            <div className="history-year-input">
                <label htmlFor="year">Year:</label>
                <input
                    id="year"
                    name="year"
                    type="number"
                    value={year}
                    min="0"
                    max="9999"
                    onChange={handleYearChange}
                />
            </div>
        </div>
    );
}