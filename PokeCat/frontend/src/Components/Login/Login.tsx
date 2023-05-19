import { Form, Input, Col, Row, Button } from "antd"
import { sha1 } from "crypto-hash"
import { addPlayer } from "../../API"
import { useState } from "react"

export interface UserData {
    username: string
    password: string
}

interface LoginProps {
    onLoginSuccess: (state: boolean) => void
}

export const Login = (props: LoginProps) => {
    const [isHidden, setIsHidden] = useState(false)

    const FormOnFinish = async (data: UserData) => {
        let userData: UserData = {username: data.username, password: ""}
        userData.password = await sha1(data.password)
        console.log(userData)
        addPlayer(userData, setIsHidden)
        props.onLoginSuccess(isHidden)
    }

    return(
        <>
            <Row justify={"center"}>
                <Col xl={6} md={12} xs={24}>
                    <Form
                        name="user"
                        onFinish={FormOnFinish}
                        labelCol={{span: 8}}
                    >
                        <Form.Item
                            name={"username"}
                            label={"Username"}
                            rules={[{required: true, message: "Please enter your username"}]}
                        >
                            <Input placeholder="Username"/>
                        </Form.Item>

                        <Form.Item
                            name={"password"}
                            label={"Password"}
                            rules={[{required: true, message: "Please enter your password"}]}
                        >
                            <Input.Password placeholder="Password"/>
                        </Form.Item>

                        <Form.Item wrapperCol={{offset: 8}}>
                            <Button type="primary" htmlType="submit">
                                Join
                            </Button>
                        </Form.Item>
                    </Form>
                </Col>
            </Row>
            
        </>
    )
}