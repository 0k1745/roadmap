"use client";


import {Input} from "antd";
import {useState} from "react";
import {sha3_512} from 'js-sha3';
import {CheckOutlined, CloseOutlined} from "@ant-design/icons";

const Baby = () => {
    const [result, setResult] = useState({state: false});

    function testValue(name: string) {
        // Step 1: Bcrypt hashing
        console.log("FIRE");
        // const saltRounds: number = 16; // You can adjust the number of salt rounds as per your requirement
        // const hashedBcrypt: string = bcrypt.hashSync(name, saltRounds);

        const hashedSha3: string = sha3_512(name); // Using SHA3-256 here, you can choose other variants if needed
        const encodedBabyName = "e93e5cf6c3ae48ce285a520e62d9a4f66090926e1e0496e5a37801b07a398969c81863e31c99064e5e71b12ffcbdf7e872a74ccac9d00de71d303529c759178b";
        console.log(`test: ${hashedSha3 === encodedBabyName}`);

        if (encodedBabyName === hashedSha3) {
            setResult({state: true})
        } else {
            setResult({state: false})
        }

    }

    function getElement() {
        if (result.state) {
            return <CheckOutlined style={{fontSize: '160px', color: "green"}}/>
        }
        return <CloseOutlined style={{fontSize: '160px', color: "red"}}/>
    }

    return <>

        <Input placeholder="Type the future name of the baby here"
               onChange={(e) => testValue(e.target.value)}
        />
        {getElement()}
    </>
}
export default Baby;