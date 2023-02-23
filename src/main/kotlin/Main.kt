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
    val user = args[0]
    val password = args[1]
    val auth = Base64.getEncoder().encodeToString("$user:$password".toByteArray())
    val url = "https://$user:$password@updates.dnsomatic.com/"

    val okHttpClient = OkHttpClient.Builder().build()
    val gson: Gson = GsonBuilder().setLenient().create()
    val retrofit = Builder().client(okHttpClient).baseUrl(url)
        .addConverterFactory(GsonConverterFactory.create(gson))
        .build()
    val apiInterface = retrofit.create(ApiInterface::class.java)

    while (true) {
        val response = apiInterface.post(auth)
        println(response.code())

        // Update IP every 30 minutes.
        delay(1800000)
    }
}

suspend fun ApiInterface.post(auth: String) = withContext(Dispatchers.IO) {
    val hostname = "all.dnsomatic.com"
    post(hostname, "Basic $auth")
}
