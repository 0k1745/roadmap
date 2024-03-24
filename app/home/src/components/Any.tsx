"use client"; // This is a client component 👈🏽
import React, { useEffect, useState } from "react";

async function funcTest(children: any): Promise<JSX.Element> {
    console.log(`Type of children is ${typeof children}`);
    switch (typeof children) {
        case "object":
            if (Array.isArray(children)) {
                return (
                    <a href={`/array/${btoa(JSON.stringify(children))}`}>
                        array
                    </a>
                );
            } else {
                return (
                    <div>
                        {Object.values(children).map((child) => {
                            return funcTest(child);
                        })}
                    </div>
                );
            }
        case "bigint":
            return <div>bigint is {children.toString()}</div>;
        case "boolean":
            return <div>boolean is {children.toString()}</div>;
        case "string":
            return <div>string is {children}</div>;
        case "number":
            return <div>number is {children}</div>;
        case "function":
            return <div>function</div>;
        case "symbol":
            return <div>symbol is {children.toString()}</div>;
        case "undefined":
            return <div>undefined is {children}</div>;
        default:
            return <div>Nope</div>;
    }
}

const Any = ({ children }: { children: any }) => {
    const [content, setContent] = useState<JSX.Element | null>(null);

    useEffect(() => {
        const getContent = async () => {
            const renderedContent = await funcTest(children);
            setContent(renderedContent);
        };

        getContent();
    }, [children]);

    return (
        <div>
            <p>{content}</p>
        </div>
    );
};

export default Any;
