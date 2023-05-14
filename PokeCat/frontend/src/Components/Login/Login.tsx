import { Form, Input, Col, Row, Button } from "antd"

export const Login = () => {

    const FormOnFinish = () => {

    }

    return(
        <>
            <Row justify={"center"}>
                <Col xl={6} md={12} xs={24}>
                    <Form
                        name="user"
                        onFinish={FormOnFinish}
                    >
                        <Form.Item
                            name={"username"}
                            rules={[{required: true, message: "Please enter your username"}]}
                        >
                            <Input placeholder="Username"/>
                        </Form.Item>

                        <Form.Item
                            name={"password"}
                            rules={[{required: true, message: "Please enter your password"}]}
                        >
                            <Input placeholder="Password"/>
                        </Form.Item>

                        <Form.Item>
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