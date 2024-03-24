"use client"; // This is a client component 👈🏽

import {Table} from "antd";

const ProcessesList = ({children}: { children: any }) => {
    return (
        <div>
            {console.log(JSON.stringify(children))}
            <Table dataSource={children.resources} columns={[
                {
                    title: 'location',
                    dataIndex: 'location',
                    key: 'location',
                    render: (text) => <a href={`${text}`}>{text}</a>
                }
            ]}></Table>
        </div>
    );
};

export default ProcessesList;
