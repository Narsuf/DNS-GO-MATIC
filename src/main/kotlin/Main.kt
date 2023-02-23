import com.google.gson.Gson
import com.google.gson.GsonBuilder
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.delay
import kotlinx.coroutines.runBlocking
import kotlinx.coroutines.withContext
import okhttp3.OkHttpClient
import retrofit2.Retrofit.Builder
import retrofit2.converter.gson.GsonConverterFactory
import java.util.*

fun main(args: Array<String>) = runBlocking {
    // Credentials and baseUrl
    val user = args[0]
    val password = args[1]
    val baseUrl = "https://$user:$password@updates.dnsomatic.com/"

    // Create ApiInterface object to make the requests.
    val okHttpClient = OkHttpClient.Builder().build()
    val gson: Gson = GsonBuilder().setLenient().create()
    val retrofit = Builder().client(okHttpClient).baseUrl(baseUrl)
        .addConverterFactory(GsonConverterFactory.create(gson))
        .build()
    val apiInterface = retrofit.create(ApiInterface::class.java)

    // Query parameter hostname and header authorization
    val hostname = "all.dnsomatic.com"
    val auth = Base64.getEncoder().encodeToString("$user:$password".toByteArray())

    // Frequency
    val minutes = 30

    while (true) {
        // Don't let losing your connection in a bad moment ruin your service.
        runCatching { apiInterface.updateIpToDdns(hostname, auth) }
            .onSuccess { println(it.code()) }
            .onFailure { println(it.message) }
            .also {
                // Update IP every x minutes.
                println("See you again in $minutes minutes")
                delay(minutes * 60000L)
            }
    }
}

suspend fun ApiInterface.updateIpToDdns(hostname: String, auth: String) = withContext(Dispatchers.IO) {
    post(hostname, "Basic $auth")
}
