"use client"; // This is a client component 👈🏽

import React, {useState, useEffect} from 'react';
import ProcessesList from "@/app/process/ProcessesList";

const Processes = () => {
    const [data, setData] = useState(null);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch('http://localhost:8080/process');
                if (!response.ok) {
                    throw new Error('Failed to fetch data');
                }
                const jsonData = await response.json();
                setData(jsonData);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        };

        fetchData();
    }, []);

    return (
        <>
            <h1 className="text-3xl font-bold underline">
                Hello world!
            </h1>
            {data ? (
                <ProcessesList>{data}</ProcessesList>
            ) : (
                <h1 className="text-3xl font-bold underline">
                    Loading...
                </h1>
            )}
        </>
    );
};

export default Processes;
