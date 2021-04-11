exports.lambdaHandler = async (event, context) => {
    const response = {
        'statusCode': 200,
        'body': JSON.stringify({
            message: 'Welcome to Pokedex API',
        })
    }

    return response;
};
