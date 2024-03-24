"use client"; // This is a client component 👈🏽
import React, {useState, useEffect} from 'react';
import {useParams} from "next/navigation";
import {Card} from "antd";

const Process = () => {
    const [data, setData] = useState(null);
    const {id} = useParams(); // Example ID state

    useEffect(() => {
        const fetchData = async (id) => {
            try {
                const response = await fetch(`http://localhost:8080/process/${id}`);
                console.log(id);
                if (!response.ok) {
                    throw new Error('Failed to fetch data');
                }
                const jsonData = await response.json();
                setData(jsonData);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        };

        fetchData(id);
    }, [id]);

    return (
        <>
            {data ? (
                    <Card title={data.Name} style={{width: 300}}>
                        <p>State {data?.State}</p>
                        <p>Memory Max {data?.Memory?.Max}</p>
                        <p>Memory Current {data?.Memory?.Current}</p>
                    </Card>
                ) : (
                <p>Loading...</p>
                )}
        </>
    );
};

export default Process;