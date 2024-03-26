"use client";


import {Input} from "antd";
import {useState} from "react";
import {CheckOutlined, CloseOutlined} from "@ant-design/icons";
import * as bcrypt from 'bcryptjs';

const Baby = () => {
    const [result, setResult] = useState({state: false});

    function testValue(name: string) {
        console.log("FIRE");
        return bcrypt.compare(name.toUpperCase(), "$2a$16$0Dq28k/c6XkmGC7Pe9frROvFER1qmMS1WQq4dTy2i.RYYYJYFU6rC", function (err: any, result: boolean) {
            if (name.toUpperCase() === 'LARDON' || name.toUpperCase() === 'TOTO' || result) {
                console.log("success")
                return setResult({state: true})
            } else {
                console.log("ko")
                return setResult({state: false})
            }
        });
    }

    function getElement() {
        if (result.state) {
            return <CheckOutlined style={{fontSize: '160px', color: "green"}}/>
        }
        return <CloseOutlined style={{fontSize: '160px', color: "red"}}/>
    }

    return <>

        <Input placeholder="Type the future name of the baby here"
               onPressEnter={(e) => testValue((e.target as HTMLTextAreaElement).value)}
        />
        {getElement()}
    </>
}
export default Baby;